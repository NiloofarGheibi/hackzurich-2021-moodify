import React from "react";
import {VictoryChart,  VictoryLine, VictoryScatter} from "victory";
import styled from '@emotion/styled';

const colors = ["#7EC8E3", "#0E86D4", "#0000FF",  "#000C66", "#050A30"];

const today = 26;

const ScatterPoint = ({ x, y, datum, min, max }) => {
    const i = React.useMemo(() => {
        return Math.floor(((datum.y - min) / (max - min)) * (colors.length - 1));
    }, [datum, min, max]);

    return <StyledPoint color={colors[i]}  cx={x} cy={y} r={4} />;
};

export default function GraphMonths(){
    const data = [
        { x: "01", y: 43 },
        { x: "02", y: 44 },
        { x: "03", y: 47 },
        { x: "04", y: 51 },
        { x: "05", y: 57 },
        { x: "06", y: 62 },
        { x: "07", y: 67 },
        { x: "08", y: 68 },
        { x: "09", y: 63 },
        { x: "10", y: 54 },
        { x: "11", y: 47 },
        { x: "12", y: 42 },
        { x: "13", y: 43 },
        { x: "14", y: 44 },
        { x: "15", y: 47 },
        { x: "16", y: 51 },
        { x: "17", y: 57 },
        { x: "18", y: 62 },
        { x: "19", y: 67 },
        { x: "20", y: 68 },
        { x: "21", y: 63 },
        { x: "22", y: 54 },
        { x: "23", y: 47 },
        { x: "24", y: 42 },
        { x: "25", y: 62 },
        { x: "26", y: 67 },
        { x: "27", y: 68 },
        { x: "28", y: 63 },
        { x: "29", y: 54 },
        { x: "30", y: 47 },
        { x: "31", y: 42 }
    ];

    const moods = data.map(({ y }) => y);
    const min = Math.min(...moods);
    const max = Math.max(...moods);

    return (
            <VictoryChart  width={1000}>
                <VictoryLine data={data}  style={{data: {stroke: "#7EC8E3"}}} />
                <VictoryScatter
                    data={data}
                    dataComponent={<ScatterPoint min={min} max={max}/>}
                />
                <VictoryLine
                    style={{
                        data: { stroke: "red", strokeWidth: 2 },
                        labels: { angle: -90, fill: "red", fontSize: 20 }
                    }}
                    x={() => today}
                />
            </VictoryChart>
    );
}

const StyledPoint = styled.circle`
  fill: ${(props) => props.color};
`;


const Container = styled.div`            
width: 80%;
`;
