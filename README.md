# Pathfinder
## What is it?
This is a small pet project ive had on my mind for a little while. Simply put its a TSP solver with a gui around it. Ill be building this out in my spare time to get a more comfortable with building out entire projects. 
pet project where users can use go, enter coordinates of sereral places they would like to visit, and a best path is returned. 

## technologies
- Golang, Javascript, React, more to come

# TODO:
- [ ] POC complete
    - [ ] users have a portal where they can enter a list of coordinates.
    - [ ] users can submit these coordinates to the server and have an optimized route returned to them
        - [x] backend support
        - [ ] front end support
    - [] basic front end styling
    - [x] basic make file added
        - [x] command to rebuild go server
        - [x] command to run node server 
    - [x] abstraction between host machine (runs in a virtual machine)
    - [x] project can be cloned and built on a clean machine (untested but as long as a machine supports vagrant there shouldnt be any issue)
    - [] nginx acts as revers proxy giving a single endpoint for users to go to

# Wishlist
- [] simulated annealing used to find route
- [] ffi interaction between rust and golang
- [] authentication for user sign on
- [] backplaneing requests - ensure that we reject additional requests instead of running the server into the ground
- [] logging to a centralized location
- [] 
- [] log queries & results for later analysis
    - [] sample queries are checked against a true solution and speed
- [] investigate hosting accross multiple machines
- [] compile react to static files so that they can be served directly by nginx 
