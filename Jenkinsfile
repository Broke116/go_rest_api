#!/usr/bin/env groovy

pipeline {
    environment {
        DOCKER_HUB_USER = 'ekinyucel'
        DOCKER_CREDENTIAL_ID = 'docker_hub'
    }

    agent none

    stages {
        stage('Pre process') {
            agent any
            steps {
                script {
                    if ("(docker ps -q -f name=go_api)") {
                        sh 'docker stop $(docker ps -a -q --filter ancestor=go_api)'
                        echo "Running container is stopped"
                        sh 'docker stop $(docker ps -a -q --filter ancestor=mongo:latest)'
                        echo "Running database container is stopped."
                    } else {
                        echo "Do not have a running container right now."
                    }
                    container_id = sh '$(docker ps | grep go_api | grep -o "^[0-9a-z]*")'
                    echo $container_id
                }                
                sh 'docker container prune'
                echo "Stopped/unused containers are pruned"
                //sh 'docker rmi -f $(docker images --format "{{.Repository}}:{{.Tag}}" | grep "go_api")'
                //echo "Existing image is removed"
            }
        }
        stage('Build') {
            agent any
            steps {
                git url: "https://github.com/Broke116/go_rest_api.git", branch: "api_compose"
                sh 'docker-compose build'
            }
        }
        stage('Run') {
            agent any
            steps {
                sh 'docker-compose up -d'
            }
        }
        stage('Test & Clean-up') {
            parallel {
                stage('Test'){
                    agent any
                    steps {
                        sh 'docker ps -q -f name=go_api'
                    }
                }
                stage('Clean up') {
                    agent any
                    steps {
                        sh 'docker rmi -f $(docker images -f "dangling=true" -q)'
                        echo "Dangling images removed"
                    }
                }
            }
        }
    }
}