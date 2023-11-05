pipeline {
    agent any

    stages {
        stage('Check Go Version') {
            steps {
                script {
                    def goVersion = sh(script: 'go version', returnStdout: true).trim()
                    echo "Go version on your local machine: $goVersion"
                }
            }
        }
    }
}


