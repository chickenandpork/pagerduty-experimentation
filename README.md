# pagerduty-oncall
Experimentation into PagerDuty API
--


In order to use this code:
1. Create an App in PagerDuty
  1.a. created a new API key, with a result such as u+tbvp8hNznmms-jbDSB
2. checking a URL such as https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1incidents/post you'll see a Test API Token
3. On a URL such as https://dev-chickenandpork.pagerduty.com/incidents choose People -> Teams to create a Team
  3.a. I created "Sample Team", tag==SRE
  3.b. on success get forwarded to a URL such as https://dev-chickenandpork.pagerduty.com/teams/P2WZKGD/users
4. On a URL such as https://dev-chickenandpork.pagerduty.com/incidents choose People -> On-Call Schedules to create a Schedule
  4.a. I created "Schedule 44", selected my "Sample Team", and added my on user to "Layer 1"
  4.b. Clicked "Create Schedule" (white on dark blue background, upper right)
  4.c. On success, I get redirected to https://dev-chickenandpork.pagerduty.com/schedules#P1NAQ02
5. On a URL such as https://dev-chickenandpork.pagerduty.com/schedules#P1NAQ02 I click https://dev-chickenandpork.pagerduty.com/escalation_policies to create an escalation policy
  5.a. New Policy
  5.b. Team: Sample Team; Notify the following: Schedule 44, save
  5.c. on success, get forwarded to a URL such as https://dev-chickenandpork.pagerduty.com/escalation_policies#PEEQ261
6. bazel run //:pdoc -- --auth_token 'u+tbvp8hNznmms-jbDSB' --schedule_id P1NAQ02
