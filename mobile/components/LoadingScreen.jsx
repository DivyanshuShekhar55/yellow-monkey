import { View, Text } from 'react-native'
import React, { useRef } from 'react'
import LottieView from 'lottie-react-native';

const LoadingScreen = () => {
    const animation = useRef(null);
    return (
        <View style={{
            backgroundColor: '#fff',
            alignItems: 'center',
            justifyContent: 'center',
            flex: 1
        }}>
            <LottieView
                autoPlay
                ref={animation}
                style={{
                    width: 200,
                    height: 400,
                    backgroundColor: '#fff',
                }}
                
                source={require('../assets/lottie/loading.json')}
            />
        </View>
    )
}

export default LoadingScreen