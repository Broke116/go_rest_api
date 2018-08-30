#!/usr/bin/env groovy

pipeline {
    environment {
        DOCKER_HUB_USER = 'ekinyucel'
        DOCKER_CREDENTIAL_ID = 'docker_hub'
    }

    agent none

    stages {
        stage('Stop running container') {
            agent any
            steps {
                sh 'docker stop $(docker ps -a -q --filter ancestor=rest_api)'
                echo "Container stopped"
                sh 'docker container prune'
                echo "Stopped containers pruned"
            }
        }
        stage('Build') {
            agent any
            steps {
                git "https://github.com/Broke116/go_rest_api.git"
                sh 'docker build -t rest_api .'
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