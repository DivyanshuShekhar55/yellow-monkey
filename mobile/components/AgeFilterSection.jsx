import { View, Text, Dimensions } from 'react-native'
import React from 'react'
import Slider from '@react-native-community/slider';

const WIDTH = Dimensions.get("screen").width

const AgeFilterSection = () => {
    return (
        <View>
            <View style={{flexDirection:"row", justifyContent:"space-between", width:WIDTH*0.9, marginBottom:12, alignItems:"center"}}>
                <Text style= {{fontWeight:"700", fontSize:24}}>Age</Text>
                <Text style={{fontWeight:"700", fontSize:16}}>20-40</Text>
            </View>
            <Slider
                style={{ width: WIDTH*0.9, height: 40 }}
                minimumValue={20}
                maximumValue={40}
                minimumTrackTintColor="#d4506b"
                maximumTrackTintColor="#000"
                step={1}
                thumbTintColor='#d45'
                
            />
        </View>
    )
}

export default AgeFilterSection