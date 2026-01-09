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
        

        stage('Go Test') {
            agent {
                docker {
                    image 'golang:1.24.11-alpine3.23'
                    args '-e GOCACHE=/tmp/go-cache'
                }
            }
            steps {
                sh '''
                  go version
                  cd go-application
                  go test ./... -coverprofile=coverage.out
                '''
            }
        }
        stage('SonarQube Ananlysis'){
            /* agent {
                docker {
                    image 'golang:1.24.11-alpine3.23'
                    args '-e GOCACHE=/tmp/gocache'
                }
            } */
            agent {
                docker {
                    image 'sonarsource/sonar-scanner-cli:latest'
                    args '-e SONAR_USER_HOME=$WORKSPACE/.sonar'  // Tell SonarScanner to store cache in Jenkins workspace
                }
            }
            steps{
                withSonarQubeEnv ('jenkins-sonar'){  // Sonar server name created in Jenkins Server
                    sh '''
                      cd go-application
                      sonar-scanner
                    '''
                }
            }
        }

        stage('Quality Gate'){
            agent any 
            steps{
                timeout(time: 5, unit: 'MINUTES'){
                    script {
                        def qg = waitForQualityGate() 
                        if (qg.status != 'OK') {
                            error "Quality Gate Failed..: ${qg.status}"
                        }
                    }
                }
            }
        }

        stage('Build Go Binary'){
            agent {
                docker {
                    image 'golang:1.24.11-alpine3.23'
                    args '-e GOCACHE=/tmp/go-cache'
                }
            }
            steps {
                sh '''
                  cd go-application
                  go mod download
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
                    ./go-application   // This the context, i,e where application source code is present
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

        /* stage('Deploy to Kubernetest'){
            agent any
            steps {
                withCredentials([file(credentialsId: 'kubeconfig-jenkins', variable: 'KUBECONFIG')]){
                    sh '''
                      set -x   # To print every command
                      set +e  # Never exit early automatically

                      NAMESPACE=default
                      DEPLOYMENT=go-db-app
                      CONTAINER=go-db-app-container
                      IMAGE="${IMAGE_NAME}:${IMAGE_TAG}"

                      echo "Step 1: Updating Deployment Image"
                      kubectl -n $NAMESPACE set image deployment/$DEPLOYMENT \
                        $CONTAINER=$IMAGE
                      SET_IMAGE_STATUS=$?

                      if [ $SET_IMAGE_STATUS -ne 0 ]; then
                        echo "Failed to set image, exiting.."
                        exit 1
                      fi

                      echo "Step 2: Waiting for rollout to complete"
                      kubectl -n $NAMESPACE rollout status deployment/$DEPLOYMENT --timeout=120s
                      ROLLOUT_STATUS=$?

                      if [ $ROLLOUT_STATUS -ne 0 ]; then
                        echo "Rollout Failed, Initiating rollback...!!"
                        kubectl -n $NAMESPACE rollout undo deployment/$DEPLOYMENT 
                        echo "Rollback Triggered"
                        exit 1
                      fi

                      echo "Deployment completed successfully..!"
                      exit 0 
                    '''
                }
            }
        } */

        }
    }
    
