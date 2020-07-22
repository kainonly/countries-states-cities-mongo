import unittest
import os
from ssl_for_qcloud import QcloudSSL


class MyTestCase(unittest.TestCase):
    def setUp(self) -> None:
        self.deploy = QcloudSSL(os.environ['qcloud_id'], os.environ['qcloud_key'])

    def test_search(self):
        response = self.deploy.search('*.kainonly.com')
        self.assertIsNotNone(response)
        print(response.TotalCount)
        print(response.Certificates)

    def test_delete(self):
        response = self.deploy.search('*.kainonly.com')
        self.assertIsNotNone(response)
        if response.TotalCount != 0:
            for v in response.Certificates:
                response = self.deploy.delete(v.CertificateId)
                self.assertIsNotNone(response)

    def test_upload(self):
        with open('../resource/fullchain.pem') as f:
            pub_key = f.read()
        with open('../resource/privkey.pem') as f:
            pri_key = f.read()
        response = self.deploy.upload(pub_key, pri_key)
        self.assertIsNotNone(response)


if __name__ == '__main__':
    unittest.main()
