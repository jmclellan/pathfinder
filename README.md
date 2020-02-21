# Pathfinder
## What is it?
This is a small pet project ive had on my mind for a little while. Simply put its a TSP solver with a gui around it. Ill be building this out in my spare time to get a more comfortable with building out enitre projects. 
pet project where users can use go, enter coordinates of sereral places they would like to visit, and a best path is returned. 

## technologies
- Golang, Javascript, React, more to come

# TODO:
- [] POC complete
    - [] users have a portal where they can enter a list of coordinates.
    - [] users can submit these coordinates to the server and have an optimized route returned to them
    - [] webpack runs and converts jsx to static files that can be run on the client
    - [] basic front end styling
    - [] basic make file added
        - [] command to rebuild go server
        - [] command to run webpack 
    - [] abstraction between host machine (either virtual machine or docker)
    - [] project can be cloned and built on a clean machine

# Wishlist
- [] simulated annealing used to find route
- [] ffi interaction between rust and golang
- [] authentication for user sign on
- [] backplannig requests - ensure that we reject additional requests instead of running the server into the ground
- [] log queries & results for later analysis
    - [] sample queries are checked against a true solution and speed
- [] investigate hosting accross multiple machines