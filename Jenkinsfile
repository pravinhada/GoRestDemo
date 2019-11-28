node {
    def goHome
    def pom

    stage('Prepare') {
        goHome = tool 'jenkins_go'
    }

    stage('Checkout') {
        checkout scm
    }

    stage('Build') {
        if(isUnix()) {
            sh "'${mvnHome}/bin/go' build"
        } else {
            bat(/"${mvnHome}\bin\go" build/)
        }
    }

    if(env.BRANCH_NAME == 'master') {
        stage('Build from master') {
            echo "Finished the master build."
        }
    }

    if(env.BRANCH_NAME == 'develop') {
        stage('Uploading Artifacts from develop') {
            echo "Finished the snapshots upload."
        }
    }
}