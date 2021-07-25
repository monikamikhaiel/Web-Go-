Vagrant.configure("2") do |config|
  # ... other configuration
  #create a vagrant box from a base image of ubuntu
  config.vm.box="ubuntu/trusty64"
  # create a bprivate network with this ip. 
  config.vm.network :"private_network", ip: "192.168.50.4"

  config.vm.provider "virtualbox" do |vb|
	## Resources for the VM 
	vb.memory = 4096
	vb.cpus = 1
  end
  ## a custom shared folder between the box and the host
  # the main shared folder is /vagrant 

  config.vm.synced_folder ".", "/vm/code/"
  #it will copy all the files in the host folder . to the folder in the machine /vm/code 
  
  # add the commands to install webserver nginx 
  config.vm.provision "shell",path:"requirements_install.sh"
  end
end
