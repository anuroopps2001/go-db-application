sudo apt install coreutils

curl -s https://packagecloud.io/install/repositories/circleci/runner/script.deb.sh?any=true | sudo bash

sudo apt-get install -y circleci-runner

export RUNNER_AUTH_TOKEN="your-runner-auth-token-here"

sudo sed -i "s/<< AUTH_TOKEN >>/$RUNNER_AUTH_TOKEN/g" /etc/circleci-runner/circleci-runner-config.yaml


sudo systemctl enable circleci-runner && sudo systemctl start circleci-runner

sudo systemctl status circleci-runner