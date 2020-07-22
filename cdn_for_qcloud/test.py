import unittest
import os
from cdn_for_qcloud import QcloudCDN


class MyTestCase(unittest.TestCase):
    def setUp(self) -> None:
        self.deploy = QcloudCDN(os.environ['qcloud_id'], os.environ['qcloud_key'])

    def test_update(self):
        response = self.deploy.update('ngx-bit.kainonly.com', 'f2QIESwT')
        self.assertIsNotNone(response)
        print(response)


if __name__ == '__main__':
    unittest.main()
