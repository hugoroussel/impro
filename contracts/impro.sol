pragma solidity >=0.4.22 <0.7.0;

// ----------------------------------------------------------------------------
// Impro contract
// ----------------------------------------------------------------------------


contract Impro {

  struct Image {
       uint timestamp;
       string perceptualHash;
       address owner;
       uint price;
   }
   
   Image[] allImages;

   //Checks if a given image is already registered. If it does returns its index, else returns -1
   function check(string memory _perceptualHash) public view returns (int exist_at){
     int index = -1;
     for (uint i = 0; i<allImages.length; i++ ){
       if(keccak256(abi.encode(allImages[i].perceptualHash)) == keccak256(abi.encode(_perceptualHash))){
         index = int(i);
       }
     }
     return index;
   }

   //Adds a new image
   function register(string memory _perceptualHash, uint price) public returns (bool success){
    uint time = now;
    bool is_registered = (check(_perceptualHash) != -1);
    if(!is_registered){
      allImages.push(Image(time, _perceptualHash, msg.sender, price));
    }
    return !is_registered;
   }

   //Buy an image
   function buy(string memory _perceptualHash) public payable {
     //the image must be registered to be bought
     int index = check(_perceptualHash);
     assert(index != -1);
     //check if the buyer sent enough for the requested price
     assert(allImages[uint(index)].price <= msg.value);
     allImages[uint(index)].owner = msg.sender;
   }

   //Transfer the ownership of your photo to another person
   function transfer(string memory _perceptualHash, address _newOwner) public {
     //Check it exists
     int index = check(_perceptualHash);
     assert(index != -1);
     //Check if the person wanting to transfer is the owner of the image
     assert(msg.sender == allImages[uint(index)].owner);
     allImages[uint(index)].owner = _newOwner;
   }

}
