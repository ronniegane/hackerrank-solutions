SELECT 
con.contest_id,
con.hacker_id,
con.name,
SUM(total_submissions),
SUM(total_accepted_submissions),
SUM(total_views),
SUM(total_unique_views)
FROM Contests con
JOIN Colleges coll ON con.contest_id = coll.contest_id
JOIN Challenges chall ON coll.college_id = chall.college_id
LEFT JOIN (SELECT challenge_id, SUM(total_views) total_views, SUM(total_unique_views) total_unique_views
FROM View_Stats 
GROUP BY challenge_id
) v ON chall.challenge_id = v.challenge_id
LEFT JOIN (SELECT challenge_id, SUM(total_submissions) total_submissions, SUM(total_accepted_submissions) total_accepted_submissions
FROM Submission_Stats 
GROUP BY challenge_id
)s ON chall.challenge_id = s.challenge_id
GROUP BY con.contest_id, con.hacker_id, con.name
HAVING (SUM(total_submissions) 
+ SUM(total_accepted_submissions) 
+ SUM(total_views) 
+ SUM(total_unique_views)) > 0;
