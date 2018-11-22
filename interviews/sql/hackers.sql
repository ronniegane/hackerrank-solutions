/*
- number of unique hackers that have made at least one submission
per day of the contest (rolling)
- hacker id and name who made the most submissions that day
*/

/* first get submissions grouped by id and date
*/
SELECT submission_date, hacker_id, COUNT(*) hacks
FROM Submissions s
GROUP BY submission_date, hacker_id;
