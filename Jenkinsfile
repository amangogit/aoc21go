pipeline{
    agent{
        docker{
            image 'golang:1.17.5-alpine'
        }
    }
    stages{
        stage('build'){
            steps{
                /usr/bin/sh 'go version'
            }
        }
    }
}