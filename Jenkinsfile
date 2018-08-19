#!/usr/bin/env groovy

pipeline {
    environment {
        DOCKER_HUB_USER = 'ekinyucel'
        DOCKER_CREDENTIAL_ID = 'docker_hub'
    }

    agent { dockerfile true }

    stages {
        stage('Unit') {
            steps {
                git "https://github.com/Broke116/go_rest_api.git"
                sh "docker image build -t rest_api"
            }
        }
    }
}