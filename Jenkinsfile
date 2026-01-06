pipeline {
    agent none

    environment {
        IMAGE_NAME = 'anuroop21/go-db-application'
        IMAGE_TAG = "${BUILD_NUMBER}"
    }

    stages {
        stage('Clean Workspace') {
            agent any
                steps {
                    cleanWs()
                }
            }
        

        stage('Checkout'){
            agent any
                steps{
                    checkout scm
                }
            }
        

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
                    -t $IMAGE_NAME:$IMAGE_TAG \
                    ./go-application
                '''
            }
        }

        stage('Push Image'){
            agent any
            steps {
                 script {
                    docker.withRegistry(
                        'https://index.docker.io/v1/',
                        'jenkins-docker-login'
                        )
                        {
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push()
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push('latest')
                    }
                 }
            }
        }

        stage('Deploy to Kubernetest'){
            agent any
            steps{
                sh '''
                  kubectl set image deployment/go-db-app \
                  go-db-app-container=${IMAGE_NAME}:${IMAGE_TAG}

                  kubectl rollout status deployment/go-db-app
                '''
            }
        }

        }
    }
