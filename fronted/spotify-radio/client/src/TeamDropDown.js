import {useState} from 'react';
import styled from "@emotion/styled";


export default function TeamDropDown(){
    const options = ["Infrastructure", "Design", "Product", "Customer Care"]
    const optionName = "TEAM"

    let [selectedOption, setSelectedOption] = useState("")

    let handleOptionChange = (e) => {
        setSelectedOption(e.target.value)
    }

    return (
        <div>
            {optionName}
            <br />
            <select onChange={handleOptionChange}>
                <option value={optionName}></option>
                {options.map((option) => <option key={option.key} value={option}>{option}</option>)}
            </select>
        </div>
    );
}


