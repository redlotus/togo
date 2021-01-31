# TOGO

My implement for the [assignment](https://github.com/manabie-com/togo/tree/v0.0.1).  
I have decided to build the one that I would like to use. Hopefully, it will grow as a side project. It is quite personal, but it is sure funnier than work on extending the provided template.   
Sorry, guys.  
This is not really a usual README file.  

## Feature

Support main features (regular user):

- Board (CRUD): category-based, this feature allows a user to group related tasks
- Tasks (CRUD): typical to-do task  
    Properties

    - Priority
    - Tags
    - Order (within board)
    - Status
- Filter and search (use built-in full-text-search from postgres - it is not that hard to migrate to es later)

Administrator:

- Access to all boards and tasks or specific one 
- Expose APIs for basics stats, dashboard, metrics:

    - total boards / tasks (done and undone)
    - created boards / tasks within a time window
    - done (not done) boards / tasks within a time window
    - avg boards / tasks per user(s) within a time window
    - number of distinct tags
    - most / least used tag
    - most / least used color 

## Planning and TODO (not really)

- [ ] Account service (I would like to separate it)
- [ ] Task Service
- [ ] Caching - gocache (instance caching) or redis / tar (centralize caching)
- [ ] Testing (okay, I want to confess that I am not good at writing tests TBH)
- [ ] Do we need task/job queue for something???
- [ ] Off-load token parsing and authorization to api gateway (nginx)
- [ ] Server Sent Events for dashboard 
- [ ] Markdown Engine 
- [ ] Build under Github Action
- [ ] Deploy on orchestration (ansible + k3s or better terraform + k3s, since this would be place on my beaglebone) (OUT OF KNOWLEDGE)
