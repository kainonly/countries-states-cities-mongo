from tencentcloud.common import credential
from tencentcloud.common.profile.client_profile import ClientProfile
from tencentcloud.common.profile.http_profile import HttpProfile
from tencentcloud.ssl.v20191205 import ssl_client, models


class QcloudSSL:
    def __init__(self, id: str, key: str):
        self.client = ssl_client.SslClient(
            credential=credential.Credential(id, key),
            region='',
            profile=ClientProfile(httpProfile=HttpProfile(endpoint='ssl.tencentcloudapi.com'))
        )

    def search(self, domain: str) -> models.DescribeCertificatesResponse:
        request = models.DescribeCertificatesRequest()
        request.SearchKey = domain
        return self.client.DescribeCertificates(request)

    def delete(self, id: str) -> models.DeleteCertificateResponse:
        request = models.DeleteCertificateRequest()
        request.CertificateId = id
        return self.client.DeleteCertificate(request)

    def upload(self, pub_key: str, pri_key: str) -> models.UploadCertificateResponse:
        request = models.UploadCertificateRequest()
        request.CertificatePublicKey = pub_key
        request.CertificatePrivateKey = pri_key
        return self.client.UploadCertificate(request)
