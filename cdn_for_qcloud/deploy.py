from tencentcloud.common import credential
from tencentcloud.common.profile.client_profile import ClientProfile
from tencentcloud.common.profile.http_profile import HttpProfile
from tencentcloud.cdn.v20180606 import cdn_client, models


class QcloudCDN:
    def __init__(self, id: str, key: str):
        self.client = cdn_client.CdnClient(
            credential=credential.Credential(id, key),
            region='',
            profile=ClientProfile(httpProfile=HttpProfile(endpoint='cdn.tencentcloudapi.com'))
        )

    def update(self, domain: str, cert_id: str) -> models.UpdateDomainConfigResponse:
        request = models.UpdateDomainConfigRequest()
        request.Domain = domain
        https = models.Https()
        https.Switch = 'on'
        cert = models.ServerCert()
        cert.CertId = cert_id
        https.CertInfo = cert
        request.Https = https
        return self.client.UpdateDomainConfig(request)
