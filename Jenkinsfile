pipeline {
    agent {
        docker {
            image 'golang:1.24.11-alpine3.23'
            args '-e GOCACHE=/tmp/go-cache'
        }
    }

    stages {
        stage('Verify Environment') {
            steps {
                sh '''
                  go version
                  whoami || true
                  cat /etc/os-release
                '''
            }
        }

        stage('Go Build & Test') {
            steps {
                sh '''
                  cd go-db-app
                  go mod download
                  go test ./...
                  go build -o app
                '''
            }
        }

        stage('Docker Build') {
            steps {
                sh '''
                  docker build \
                    -f docker/docker/Dockerfile \
                    -t go-db-app:${BUILD_NUMBER} \
                    go-db-app
                '''
            }
        }
    }
}
