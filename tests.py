"""
    Workflow for exceptions TestCase
    Assert if the methods are dispatching the correct error codes
"""
import unittest
from app import app


class TestWorkflowSimple(unittest.TestCase):
    client = app.test_client()
    def test_post_400(self):
        """
            Test if the body is a json valid
        """
        self.assertEqual(self.client.post('/workflow/', data='wrong json').status_code, 400)

    def test_post_403(self):
        """
            Test if status passed in body are between inserted or consumed
        """
        data = {'status': 'inexistent status'}
        self.assertEqual(self.client.post('/workflow/', json=data).status_code, 403)

    def test_post_200(self):
        data = {
            "data": {
                "field1": "value1",
                "child1": {
                    "field1": "value1"
                    }
                },
                "child2": {
                    "field1": "value1"
                },
            "status": "inserted",
            "steps": ["step1", "step2", "step3"]
        }
        self.assertEqual(self.client.post('/workflow/', json=data).status_code, 201)


if __name__ == '__main__':
    unittest.main()
