node {
    stage('kxcab') {
        if (TAG == 'kxcab') {
            echo '开始部署[kxcab]'
            sshagent (credentials: ['Jenkins']) {
                sh "ssh -o StrictHostKeyChecking=no root@39.106.77.239 '${deploy(TAG)}'"
            }
        }
    }
    stage('kxcab-dev') {
        if (TAG == 'kxcab-dev') {
            echo '开始部署[kxcab-dev]'
            sshagent (credentials: ['Jenkins']) {
                sh "ssh -o StrictHostKeyChecking=no root@39.106.77.239 '${deploy(TAG)}'"
            }
        }
    }
    stage('ydcab') {
        if (TAG == 'ydcab') {
            echo '开始部署[ydcab]'
            sshagent (credentials: ['Jenkins']) {
                sh "ssh -o StrictHostKeyChecking=no root@39.106.77.239 '${deploy(TAG)}'"
            }
        }
    }
    stage('ydcab-dev') {
        if (TAG == 'ydcab-dev') {
            echo '开始部署[ydcab-dev]'
            sshagent (credentials: ['Jenkins']) {
                sh "ssh -o StrictHostKeyChecking=no root@39.106.77.239 '${deploy(TAG)}'"
            }
        }
    }
    stage('tbcab') {
        if (TAG == 'tbcab') {
            echo '开始部署[tbcab]'
            sshagent (credentials: ['Jenkins']) {
                sh "ssh -o StrictHostKeyChecking=no root@39.106.77.239 '${deploy(TAG)}'"
            }
        }
    }
    stage('tbcab-dev') {
        if (TAG == 'tbcab-dev') {
            echo '开始部署[tbcab-dev]'
            sshagent (credentials: ['Jenkins']) {
                sh "ssh -o StrictHostKeyChecking=no root@39.106.77.239 '${deploy(TAG)}'"
            }
        }
    }
}

def deploy(tag) {
    def url = "https://${tag}.auroraride.com/maintain/update/iemANTrAplaSTuRAMetBAHureAVaTertRiUMShrOWpUraNCfaseNtIderIANsGUE"
    def str = """
        docker pull registry-vpc.cn-beijing.aliyuncs.com/liasica/cabservd:$tag
        curl $url
        docker stop ${tag}
        docker rm -f ${tag}
        mkdir -p /var/www/cabservd/${tag}/runtime
        docker run -itd --name ${tag} --restart=always \
        --network host \
        -v /var/www/cabservd/${tag}/config:/app/config \
        -v /var/www/cabservd/${tag}/runtime:/app/runtime \
        registry-vpc.cn-beijing.aliyuncs.com/liasica/cabservd:$tag
        docker image prune -f
        docker container prune -f
        docker volume prune -f
    """
    return str
}