# Pathfinder
## What is it?
This is a small pet project ive had on my mind for a little while. Simply put its a TSP solver with a gui around it. Ill be building this out in my spare time to get a more comfortable with building out entire projects. 
pet project where users can use go, enter coordinates of sereral places they would like to visit, and a best path is returned. 

## technologies
- Golang, Javascript, React, more to come

# TODO:
- [ ] POC complete
    - [x] users have a portal where they can enter a list of coordinates.
    - [ ] users can submit these coordinates to the server and have an optimized route returned to them
        - [x] backend support
        - [ ] front end support
        - [ ] communication between the two via reverse proxy for single enpoint
    - [x] basic front end styling
    - [x] basic make file added
        - [x] command to rebuild go server
        - [x] command to run node server 
    - [x] abstraction between host machine (runs in a virtual machine)
    - [x] project can be cloned and built on a clean machine (untested but as long as a machine supports vagrant there shouldnt be any issue)

# Wishlist
- [ ] optimize path finding
    - [ ] decrease number of permutations searched to cut search space in half
    - [ ] add nieve random searching to allow for solution to more end points
    - [ ] add simulated annealing find best path
    - [ ] add controls to be able to commit a specific amount of computation time to calculating an optimal route
    - [ ] use go ffi to leverage speed and efficiency of a rust library
- [ ] integrations
    - [ ] add user accounts & authentication
    - [ ] add database to periodically store user trips for analysis later
        - [ ] explore storing summary data
        - [ ] store longer data set which can be used to test proposed solution against true solution and measured for speed
- [ ] add centralized logging
    - [ ] log rotate set up upon vagrant up
    - [ ] logging levels used
- [ ] improve system resiliancy
    - [ ] Dockerize/Pod-ify application
        - [ ] use system to manage clusters
    - [ ] application on vm to restart services as required
    - [ ] do not allow system to get overloaded when optimizing requests
        - [ ] backplaneing returns appropriate error code instead of taking on too much work
        - [ ] queue requests if they cannot currently be run
    - [ ] investigate hosting application over multiple machines
- [ ] optimize website
    - [ ] profile react componants
    - [ ] compile react code to static js files and serve directly via nginx or apache
    - [ ] profile reverse proxy
- [ ] front end improvements
    - [ ] add visualization for coordinates
    - [ ] add map
    - [ ] add visualization during optimization (connect all points and randomly iterate through paths)
- [ ] functionality improvements
    - [ ] users can select a starting point from which all paths begin
    - [ ] users can select an ending point at which all paths end (potentially the starting point)
    - [ ] users users can see their search history
    - [ ] users can look up famous lanmarks by name and their long & lat will be added automatically
    - [ ] some sort of admin portal where we can see uptodate monitoring of the system
