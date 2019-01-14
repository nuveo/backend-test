"""
    Workflow for exceptions TestCase
    Assert if the methods are dispatching the correct error codes
"""
import unittest
from app import app


class TestWorkflowSimple(unittest.TestCase):
    def test_post_400(self):
        """
            Test if the body is a json valid
        """
        client = app.test_client()
        self.assertEqual(client.post('/workflow/', data='wrong json').status_code, 400)

    def test_post_403(self):
        """
            Test if status passed in body are between inserted or consumed
        """
        client = app.test_client()
        data = {'status': 'inexistent status'}
        self.assertEqual(client.post('/workflow/', json=data).status_code, 403)


if __name__ == '__main__':
    unittest.main()
