import React, {useState} from "react";
import styled from '@emotion/styled';
import GraphMonths from "./GraphMonths";
import LocationDropDown from "./LocationDropDown";
import Filters from "./Filters";
import TimeDropDown from "./TimeDropDown";
import TeamDropDown from "./TeamDropDown";
import OverallMoodChart from "./OverallMoodChart";



export default function AdminView(){
    const actions = [
        { todo: "Setup a 1-1 with your team members", when: "sad" },
        { todo: "What about a team event? To spice things up", when: "calm" },
        { todo: "Keep up the good work and the vibe!", when: "energetic" },
        { todo: "Celebrate small things, it's important to show your team member's value", when: "happy" }
    ];

    return (
        <GlobalContainer>
        <Container>
            <ActionsBar>
                <h1 style={{color:'#808080'}}>What can we do?</h1>
                    {actions.map((action) => (
                        <ActionItem key={action.id}>{action.todo}</ActionItem>
                    ))}
            </ActionsBar>
            <Graphs>
                <Filters />
                <OverallMoodChart />
            </Graphs>
        </Container>
            <styledDiv>
                <GraphMonths />
            </styledDiv>

        </GlobalContainer>
    );
}

const styledDiv = styled.div`

  width: 70%;
`;
const GlobalContainer = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  position: relative;
`;

const Container = styled.div`
  display: flex;
  flex-direction: raw;
  width: 100%;
  height: 100%;
`;



const ActionsBar = styled.div`
  display: flex;
  flex-direction: column;
  width: 500px;
  // border-right: 4px solid #ddd;
  padding: 10px 20px;
  // background-color: #ddd;
`;

const ActionItem = styled.button`
  background-color: #DCDCDC;
  border-color: #DCDCDC;
  padding: 15px 15px 15px 15px;
  border-radius: 8px;
  margin-top: 10px;
  border-top: 1px solid rgba(0, 0, 0, 0.1);
  font-size: 22px;
   &&:hover {
  text-decoration: underline;
}
`;

const Graphs = styled.div`
width: 100%;
height: 100%;
position: relative 
`;
//
// display:flex;
// flex-direction: column;
// justify-content: center;
// align-items: center;
// margin-top: 20px;
// margin-bottom: 20px;
// width: 100%;
// height: 100%;
