// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract MyContract {
    uint256 public myUintValue;
    address[] public addressArray;
    mapping(uint256 => uint256) public myMapping;

    event ValueUpdated(uint256 newValue, address sender);

    constructor(uint256 _initialValue) {
        myUintValue = _initialValue;
    }

    function setUintValue(uint256 _newValue) public {
        myUintValue = _newValue;
        emit ValueUpdated(myUintValue, msg.sender);
    }

    function addAddressToArray(address _newAddress) public {
        addressArray.push(_newAddress);
    }

    function removeLastAddressFromArray() public {
        require(addressArray.length > 0, "Array is empty");
        addressArray.pop(); // Removes the last element
    }

    function addMultipleValuesToMapping(
        uint256[] calldata _keys,
        uint256[] calldata _values
    ) public {
        require(
            _keys.length == _values.length,
            "Keys and values must have same length"
        );
        for (uint256 i = 0; i < _keys.length; i++) {
            myMapping[_keys[i]] = _values[i];
        }
    }

    function performComplexCalculation(uint256 _a, uint256 _b) public {
        uint256 tempResult = (_a * _b) + (myUintValue / 2);
        if (tempResult > 1000) {
            myUintValue = tempResult - 500;
        } else {
            myUintValue = tempResult + 100;
        }
    }

    event LogMessage(string message);

    function triggerEventAndReturn(
        string memory _message
    ) public returns (bool) {
        emit LogMessage(_message);
        return true;
    }

    function conditionalExecution(bool _shouldRun) public {
        require(msg.sender != address(0), "Invalid sender"); // Basic require check

        if (_shouldRun) {
            myUintValue += 1;
        } else {
            myUintValue += 0;
        }
    }

    function getUintValue() public view returns (uint256) {
        return myUintValue;
    }

    function getMappingValue(uint256 _key) public view returns (uint256) {
        return myMapping[_key];
    }

    function getAddressArrayLength() public view returns (uint256) {
        return addressArray.length;
    }

    function getAddressArrayElement(
        uint256 _index
    ) public view returns (address) {
        return addressArray[_index];
    }
}
