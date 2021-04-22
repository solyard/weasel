pipeline {
    agent any
    stages {
        stage('Build Docker Image') {
            steps {
                sh "docker build . -f Dockefile -t weasel:develop"
            }
        }
    }
}
