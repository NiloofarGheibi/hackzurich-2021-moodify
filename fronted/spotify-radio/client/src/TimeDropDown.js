import {useState} from 'react';


export default function TimeDropDown(){
    const options = ["Weekly", "Daily", "Monthly"]
    const optionName = "TIME"

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