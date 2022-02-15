pipeline {
    agent any
    stages {
        stage('Build Docker Image') {
            steps {
                sh "docker build . -f Dockerfile -t weasel:develop"
                sh "echo 'test'"
            }
        }
    }
}
