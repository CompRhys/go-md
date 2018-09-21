# moldyn

This is a simple molecular dynamics package built in go.

The idea behind this work is to build up a package that is highly readable allowing for students to get an easy idea of how different algorithms are implemented. In full scale HPC packages the use of MPI structures, domain decomposition and other parallel computing methods can make the code difficult to understand for begineers.

However in order to achieve improved performance we will make use of simple paralleisation using channels in-built to go. These let us run separate processes from the same main program enabling simple speed up to be achieved for concurrent tasks.
