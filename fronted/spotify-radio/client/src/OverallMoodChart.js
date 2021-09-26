import React from 'react';
import ReactDOM from 'react-dom';
import { VictoryTheme, VictoryChart, VictoryAxis, VictoryArea, VictoryLabel, VictoryLegend } from 'victory';
import styled from "@emotion/styled";

const data_energetic = [
    { x: 0, y: 6 },
    { x: 6, y: 6 }
];

const data_sad = [
    { x: 0, y: -6 },
    { x: 7, y: -6 }
];

const data_calm = [
    { x: 0, y: 4 },
    { x: -5, y: 4 }
];

const data_happy = [
    { x: 0, y: -2 },
    { x: -2, y: -2 }
];

const Container = styled.div`
   position: absolute;              
   left: 50%;                         
   transform: translateX(-50%);
   width: 65%;
`;


export default function OverallMoodChart(){
        return (
            <Container>
            <VictoryChart
                theme={VictoryTheme.grayscale}
            >
                <VictoryAxis crossAxis
                             width={400}
                             height={400}
                             domain={[-10, 10]}
                             theme={VictoryTheme.material}
                             standalone={false}
                />
                <VictoryAxis dependentAxis crossAxis
                             width={400}
                             height={400}
                             domain={[-10, 10]}
                             theme={VictoryTheme.material}
                             standalone={false}
                />
                <VictoryLabel
                    x={200}
                    y={32}
                    text="Energetic"
                />
                <VictoryArea
                    style={{data: { fill: 'DarkOrange'}}}
                    data={data_energetic}
                />
                <VictoryLabel
                    x={415}
                    y={150}
                    text="Sad"
                />
                <VictoryArea
                    style={{ data: { fill: 'SteelBlue' } }}
                    data={data_sad}
                />
                <VictoryLabel
                    x={205}
                    y={270}
                    text="Calm"
                />
                <VictoryArea
                    style={{ data: { fill: "MediumVioletRed" } }}
                    data={data_calm}
                />
                <VictoryLabel
                    x={8}
                    y={150}
                    text="Happy"
                />
                <VictoryArea
                    style={{ data: { fill: "SkyBlue" } }}
                    data={data_happy}
                />
                <VictoryLegend x={0} y={0}
                               title="Mood indications"
                               gutter={18}
                               orientation="vertical"
                               style={{ border: { stroke: "black" }, title: { fontSize: 6 }}}
                               data={[
                                   { name: "Anxious/Frantic", symbol: { fill: "DarkOrange" } },
                                   { name: "Exuberance", symbol: { fill: "MediumVioletRed" } },
                                   { name: "Depression", symbol: { fill: "SteelBlue" } },
                                   { name: "Contentment", symbol: { fill: "SkyBlue" } }
                               ]}
                />
            </VictoryChart>
            </Container>
        )
}