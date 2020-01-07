#!/usr/bin/python3
import subprocess
import configparser


class Config(configparser.ConfigParser):
    def optionxform(self, optionstr):
        return optionstr


try:
    # Uninstall old versions
    subprocess.run('sudo apt -y remove docker docker-engine docker.io containerd runc', shell=True)
    # Install using the repository
    subprocess.run('sudo apt update', shell=True)
    # Install packages to allow apt to use a repository over HTTPS
    subprocess.run('sudo apt -y install \
      apt-transport-https \
      ca-certificates \
      curl \
      gnupg2 \
      software-properties-common', shell=True)
    # Add Docker’s official GPG key
    subprocess.run('curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -', shell=True)
    # Verify that you now have the key with the fingerprint
    subprocess.run('sudo apt-key fingerprint 0EBFCD88', shell=True)
    # Use the following command to set up the stable repository
    subprocess.run('sudo add-apt-repository \
      "deb [arch=amd64] https://download.docker.com/linux/debian \
       $(lsb_release -cs) \
       stable"', shell=True)
    # Update the apt package index
    subprocess.run('sudo apt update', shell=True)
    # Install the latest version of Docker Engine
    subprocess.run('sudo apt -y install docker-ce docker-ce-cli containerd.io', shell=True)
    # Enable Docker Service
    subprocess.run(['systemctl', 'enable', 'docker'])
    # Run this command to download the current stable release of Docker Compose
    subprocess.run('sudo curl -L \
     "https://github.com/docker/compose/releases/download/1.25.0/docker-compose-$(uname -s)-$(uname -m)" \
     -o /usr/local/bin/docker-compose', shell=True)
    # Apply executable permissions to the binary
    subprocess.run('sudo chmod +x /usr/local/bin/docker-compose', shell=True)

    file_service = '/lib/systemd/system/docker.service'
    config = Config()
    config.read(file_service)
    config.set(
        'Service',
        'ExecStart',
        '/usr/bin/dockerd -H tcp://127.0.0.1:2375 -H fd:// --containerd=/run/containerd/containerd.sock'
    )

    with open(file_service, mode='w') as f:
        config.write(f)

    subprocess.run(['systemctl', 'daemon-reload'])
    subprocess.run(['systemctl', 'restart', 'docker'])
    print('docker successfully deployed!')
except Exception as error:
    print(error)
