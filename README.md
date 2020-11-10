# cqrsGolang

<h2>Welcome to the exercises for CQRS in GoLang.  </h2>


Once you have checked out the code, you will need the following set up on your workstation:
1. A local installation of EventStoreDB from  https://www.eventstore.com/ is required to run this project.  Make sure 
that the event store database is installed and running.  This can be verified by
loading http://localhost:2113/web/index.html#/ The default username/password is admin/changeit.  Commands cannot be issued 
to the hotel reservation system without an active event store.  There are 2 methods of installing event store.
    1. A docker-compose.yml file has also been provided in the repositories root directory. Execute the command
    `docker-compose up` in the root directory if you have docker installed.
    1. You can find installation instructions for Event Store Db at https://www.eventstore.com/. 


