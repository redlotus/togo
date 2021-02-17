# TOGO

Initially, this is my implementation for the [assessment](https://github.com/manabie-com/togo/tree/v0.0.1).  However, I could not keep it up and dropped halfway.
Hence, this becomes my side project where I can recall my knowledge and do some experiments.  
This is not really a usual README file.  

## Feature

Support main features (regular user):

- Board (CRUD): category-based, this feature allows a user to group related tasks
- Tasks (CRUD): typical to-do task  
    Properties

    - Priority Levels
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

- [ ] PostgreSQL as backend (for fun, since I like to write a lot of SQL for data processing instead of ORM or Query Builder style). We can cut down the latency too.
- [ ] Account service (I would like to separate it)
- [ ] Task Service
- [ ] Caching - gocache (instance caching) or redis / tar (centralize caching)
- [ ] Testing (okay, I want to confess that I am not good at writing tests, TBH)
- [ ] Do we need a task/job queue for something???
- [ ] Off-load token parsing and authorization to api gateway (nginx)
- [ ] Server Sent Events for dashboard 
- [ ] Markdown Engine 
- [ ] Build under Github Action
- [ ] Deploy on orchestration (ansible + k3s or better terraform + k3s, since I would place this on my Beaglebone board) (OUT OF KNOWLEDGE atm)

## ON THE WAY (or references)

- https://www.clever-cloud.com/blog/engineering/2015/05/20/why-auto-increment-is-a-terrible-idea/ < old article I dig out of my stack 
- More about database problem:  https://stackoverflow.com/a/42217872
- And more: https://stackoverflow.com/a/40325406
- Re-read (3rd time maybe) SQL Performance Explained book
- Re-read Kubernetes: Up and Running: Dive Into the Future of Infrastructure
- Read The Kubernetes Book
- https://rakyll.org/style-packages/
- https://jensrantil.github.io/post/salt-vs-ansible/ < I decided re-learn Saltstack
- https://threedots.tech/tags/building-business-applications/ < outstanding articles on building maintainable service in go, highly recommended
- https://instil.co/blog/static-vs-dynamic-types/ < this and the below two were fun to read
- https://instil.co/blog/static-types-wont-save-us-from-bad-code/
- https://instil.co/blog/why-strong-static-typing-is-your-friend/
