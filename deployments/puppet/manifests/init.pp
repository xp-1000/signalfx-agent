# Main class that installs and configures the agent
class signalfx_agent (
  $config                 = lookup('signalfx_agent::config', Hash, 'deep'),
  $package_stage          = 'release',
  $repo_base              = 'splunk.jfrog.io/splunk',
  $config_file_path       = $::osfamily ? {
    'debian'  => '/etc/signalfx/agent.yaml',
    'redhat'  => '/etc/signalfx/agent.yaml',
    'windows' => 'C:\\ProgramData\\SignalFxAgent\\agent.yaml',
    default   => '/etc/signalfx/agent.yaml'
  },
  $agent_version          = '',
  $package_version        = '',
  $installation_directory = 'C:\\Program Files\\SignalFx',
  $service_user           = 'signalfx-agent',  # linux only
  $service_group          = 'signalfx-agent',  # linux only
  $apt_gpg_key            = 'https://splunk.jfrog.io/splunk/signalfx-agent-deb/splunk-B3CD4420.gpg',
  $yum_gpg_key            = 'https://splunk.jfrog.io/splunk/signalfx-agent-rpm/splunk-B3CD4420.pub',
  $manage_repo            = true  # linux only
) {

  $service_name = 'signalfx-agent'

  if !$config['signalFxAccessToken'] {
    fail("The \$config parameter must contain a 'signalFxAccessToken'")
  }

  if $::osfamily == 'windows' {
    if $agent_version == '' {
      fail("The \$agent_version parameter must be set to a valid SignalFx Agent version")
    }

    $split_config_file_path = $config_file_path.split('\\\\')
    $config_parent_directory_path = $split_config_file_path[0, - 2].join('\\')
  } else {
    $split_config_file_path = $config_file_path.split('/')
    $config_parent_directory_path = $split_config_file_path[0, - 2].join('/')

    if $package_version == '' {
      if $agent_version == '' {
        $version = 'latest'
      } else {
        $version = "${agent_version}-1"
      }
    } else {
      $version = $package_version
    }
  }

  case $::osfamily {
    'debian': {
      class { 'signalfx_agent::debian_repo':
        repo_base     => $repo_base,
        package_stage => $package_stage,
        apt_gpg_key   => $apt_gpg_key,
        manage_repo   => $manage_repo,
      }
      -> package { $service_name:
        ensure => $version,
      }
    }
    'redhat': {
      class { 'signalfx_agent::yum_repo':
        repo_base     => $repo_base,
        package_stage => $package_stage,
        yum_gpg_key   => $yum_gpg_key,
        manage_repo   => $manage_repo,
      }
      -> package { $service_name:
        ensure => $version,
      }
    }
    'windows': {
      class { 'signalfx_agent::win_repo':
        repo_base              => 'dl.signalfx.com',
        package_stage          => $package_stage,
        version                => $agent_version,
        installation_directory => $installation_directory,
        service_name           => $service_name,
        config_file_path       => $config_file_path,
      }
    }
    default: {
      fail("Your OS (${::osfamily}) is not supported by the SignalFx Agent")
    }
  }

  -> service { $service_name:
    ensure => true,
    enable => true,
  }

  if $::osfamily != 'windows' {
    class { 'signalfx_agent::service_owner':
      service_name  => $service_name,
      service_user  => $service_user,
      service_group => $service_group,
    }
  }

  file { $config_parent_directory_path:
    ensure => 'directory',
  }

  file { $config_file_path:
    ensure  => 'file',
    content => template('signalfx_agent/agent.yaml.erb'),
    mode    => '0600',
  }

  File[$config_parent_directory_path] ~> File[$config_file_path] ~> Service[$service_name]
}
