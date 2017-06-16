pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh "docker build -t campbel/myapp ."
            }
        }
        stage('Push') {
            steps {
                sh "docker push campbel/myapp"
            }
        }
    }
}