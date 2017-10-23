# one-piece-reminder
personal toy

the idea is to remind of the latest comic update from http://www.onepiece.cc/comic/ .
My first thought was to notify me on computer,for me,I have thought about the notify-send widget pretty handy on Ubuntu.But I gave up that idea because I want to make sure I won't miss any notification and notify-send don't have manual operation.

So I turn to the Wechat platform.First,I setup my server on Aliyun and test if it can provide web service,then I follow the instruction on wechat developer platform to verify my server,things was going well at that time.Almost there...I felt excited.

When I nearly finished my code,to my disappointment, I found 'message broadcasting' API was not avaliable for the personal account.So the core ```func SendNotification(issueNo, title, link string)```cannot be implemented.

For now, I turn to some SMS service provider trying to figure out if I can use SMS to notify me...But these providers are not friendly for personal user...

To be continued...

