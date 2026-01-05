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
                    -f Docker/Dockerfile \
                    -t go-db-app:${BUILD_NUMBER} \
                    go-db-app
                '''
            }
        stage('Push Image'){
            agent any
            step {
                 script {
                    docker.withRegistry('','jenkins-docker-login'){
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push()
                        docker.image("${IMAGE_NAME}:latest").push()
                    }
                 }
            }
        }
        stage('Deploy to Kubernetest'){
            agent any
            steps{
                sh '''
                  kubectl set image deployment/go-db-app \
                  app=${IMAGE_NAME}:${IMAGE_TAG}

                  kubectl rollout status deployment/go-db-app
                '''
            }
        }

        }
    }
}
