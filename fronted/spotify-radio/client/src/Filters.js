import React, {useState} from 'react';
import './App.css';
import LocationDropDown from "./LocationDropDown";
import TimeDropDown from "./TimeDropDown";
import TeamDropDown from "./TeamDropDown";
import styled from "@emotion/styled";

export default function Filters(){
    return(
        <Container>
            <LocationDropDown />
            <TimeDropDown />
            <TeamDropDown />
        </Container>
    )
}

const Container = styled.div`
  display: flex;
  flex-direction: raw;
  justify-content: space-around;
  width: 100%;
  margin-top: 10px;
  margin-bottom: 20px;
`;