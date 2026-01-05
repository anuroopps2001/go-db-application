pipeline {
    agent none

    stages {

        stage('Go Build & Test') {
            agent {
                docker {
                    image 'golang:1.24.11-alpine3.23'
                    args '-e GOCACHE=/tmp/go-cache'
                }
            }
            steps {
                sh '''
                  go version
                  cd go-db-app
                  go mod download
                  go test ./...
                  go build -o app
                '''
            }
        }

        stage('Docker Build') {
            agent any
            steps {
                sh '''
                  docker version
                  docker build \
                    -f docker/docker/Dockerfile \
                    -t go-db-app:${BUILD_NUMBER} \
                    go-db-app
                '''
            }
        }
    }
}
