pipeline{
    agent{
        docker{
            image 'golang:1.14'
        }
    }
    stages{
        stage('Test'){
            steps{
               sh 'go version' 
            }
        }
    }
}