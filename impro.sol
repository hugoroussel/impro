pragma solidity >=0.4.22 <0.7.0;

// ----------------------------------------------------------------------------
// Impro contract
// ----------------------------------------------------------------------------


contract Impro {

  struct Image {
       uint timestamp;
       address payable owner;
       int price;
   }

   mapping (string => Image) images;
   mapping (string => uint) exist;

   //Timestamp and registers the new image given its perceptual hash
   function register(string memory _perceptualHash, int price) public{
    uint time = now;
    //check if someone has already uploaded the same hash
    assert(exist[_perceptualHash] == 0);
    //if it doesn't adds it to the mapping
    images[_perceptualHash] = Image(time,  msg.sender, price);
    //add the hash to the mapping of set images
    exist[_perceptualHash] = 1;
   }

   //Checks if a given image is already registered. If it does returns true
   function exists(string memory _perceptualHash) public view returns (bool){
     return (exist[_perceptualHash] == 1);
   }

   //Buy the ownership of an image
   function buy(string memory _perceptualHash) public payable{
     address payable oldOwner = images[_perceptualHash].owner;
     address payable newOwner = msg.sender;
     //Image must exists
     assert(exists(_perceptualHash));
     //Owner can't buy hiw own image
     assert(msg.sender != getOwner(_perceptualHash));
     //Sufficient price is paid
     assert(int(msg.value) >= getPrice(_perceptualHash));
     //Change the owner
     images[_perceptualHash].owner = newOwner;
     //When the image is bought, the price is set to non-buyable
     images[_perceptualHash].price = -1;
     //Send the old owner the price of the image
     oldOwner.transfer(msg.value);
   }

   //Transfer the ownership of an image
   function transfer(string memory _perceptualHash, address payable _newOwner) public payable{
     //Image must exists
     assert(exists(_perceptualHash));
     //Check person doing the transaction is owner of image
     assert(msg.sender == getOwner(_perceptualHash));
     //Change the owner
     images[_perceptualHash].owner = _newOwner;
     //When the image is transferred, the price is set to non-buyable
     images[_perceptualHash].price = -1;
   }

   //Gets timestamp of an image
   function getTimestamp(string memory _perceptualHash) public view returns (uint price){
     assert(exists(_perceptualHash));
     return images[_perceptualHash].timestamp;
   }

   //Gets price of an image
   function getPrice(string memory _perceptualHash) public view returns (int price){
     assert(exists(_perceptualHash));
     return images[_perceptualHash].price;
   }

   //Gets owner of an image
   function getOwner (string memory _perceptualHash) public view returns (address owner){
     assert(exists(_perceptualHash));
     return images[_perceptualHash].owner;
   }


   function changePrice(string memory _perceptualHash, int _price) public {
     assert(exists(_perceptualHash));
     assert(msg.sender == getOwner(_perceptualHash));
     images[_perceptualHash].price = _price;
   }
}
