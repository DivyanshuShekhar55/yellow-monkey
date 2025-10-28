import { View, Text } from 'react-native'
import React from 'react'
import LottieView from 'lottie-react-native';

const LoadingScreen = () => {
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
                    height: 200,
                    backgroundColor: '#eee',
                }}
                // Find more Lottie files at https://lottiefiles.com/featured
                source={require('../assets/lottie/loading.json')}
            />
        </View>
    )
}

export default LoadingScreen