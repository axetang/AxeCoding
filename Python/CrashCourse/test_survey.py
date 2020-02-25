import unittest
from survey import AnonymousSurvey


class TestAnonymousSurvey(unittest.TestCase):
    def setUp(self):
        self.question = "What language did you first learn to speak?"
        self.my_survey = AnonymousSurvey(self.question)
        self.responses = ['English', 'Spanish', 'Mandarin']

    def test_store_single_response(self):
        self.my_survey.store_reponse(self.responses[0])
        self.assertIn(self.responses[0], self.my_survey.responses)

    def test_store_three_responses(self):
        for response in self.responses:
            self.my_survey.store_reponse(response)
        for response in self.responses:
            self.assertIn(response, self.my_survey.responses)

    def test_show_question(self):
        self.assertIn(self.question, self.my_survey.question)


unittest.main()
