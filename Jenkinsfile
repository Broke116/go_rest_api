#!/usr/bin/env groovy

pipeline {
    environment {
        DOCKER_HUB_USER = 'ekinyucel'
        DOCKER_CREDENTIAL_ID = 'docker_hub'
    }

    agent none

    stages {
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