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
                    if [["docker ps -q -f name=rest_api" ]] then
                        sh 'docker stop $(docker ps -a -q --filter ancestor=rest_api)'
                        echo "Running container is stopped"
                    fi
                }                
                sh 'docker container prune'
                echo "Stopped/unused containers are pruned"
                sh 'docker rmi -f $(docker images --format "{{.Repository}}:{{.Tag}}" | grep "rest_api")'
                echo "Existing image is removed"
            }
        }
        stage('Build and Run') {
            agent any
            steps {
                git url: "https://github.com/Broke116/go_rest_api.git", branch: "api_compose"
                //sh 'docker build -t rest_api .'
                sh 'docker-compose up -d'
            }
        }
        stage('Docker Run'){
            agent any
            steps {
                sh 'docker run -d --rm -p 4000:3030 -t rest_api'
                echo "Application started on port: 4000"
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