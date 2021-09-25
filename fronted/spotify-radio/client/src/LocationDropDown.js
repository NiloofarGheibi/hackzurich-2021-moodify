import {useState} from 'react';


export default function LocationDropDown(){
    const options = ["Home", "Office"]
    const optionName = "LOCATION"

    let [selectedOption, setSelectedOption] = useState("")

    let handleOptionChange = (e) => {
        setSelectedOption(e.target.value)
    }

    return (
        <div>
            {optionName}
            <br />
            <select onChange={handleOptionChange} style={{width: "100%" }}>
                <option value={optionName}></option>
                {options.map((option) => <option key={option.key} value={option}>{option}</option>)}
            </select>
        </div>
    );
}