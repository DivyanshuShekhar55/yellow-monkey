import React, { useState } from 'react';
import {
    View,
    Text,
    Image,
    StyleSheet,
    SafeAreaView,
    StatusBar,
    TouchableOpacity,
    ScrollView,
    Animated,
    PanGestureHandler,
    State,
} from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { useLocalSearchParams, useRouter } from "expo-router"

export default function GroupDetailScreen({ route }) {
    const router = useRouter();
    const { group_str } = useLocalSearchParams()
    const group = group_str ? JSON.parse(group_str) : {}
    console.log(group.image)
    const [currentStatus, setCurrentStatus] = useState(group.status);

    const getStatusConfig = (status) => {
        switch (status) {
            case 'join':
                return {
                    text: 'Join',
                    icon: 'add-outline',
                    color: '#4CAF50',
                    bgColor: '#4CAF50',
                };
            case 'requested':
                return {
                    text: 'Requested',
                    icon: 'time-outline',
                    color: '#FF9800',
                    bgColor: '#FF9800',
                };
            case 'open':
                return {
                    text: 'Open',
                    icon: 'checkmark-outline',
                    color: '#2196F3',
                    bgColor: '#2196F3',
                };
            default:
                return {
                    text: 'Join',
                    icon: 'add-outline',
                    color: '#4CAF50',
                    bgColor: '#4CAF50',
                };
        }
    };

    const handleStatusPress = () => {
        if (currentStatus === 'join') {
            setCurrentStatus('requested');
        } else if (currentStatus === 'requested') {
            setCurrentStatus('open');
        }
    };

    const SwipeToJoinButton = () => {
        const translateX = new Animated.Value(0);
        const [isCompleted, setIsCompleted] = useState(false);

        const onGestureEvent = Animated.event(
            [{ nativeEvent: { translationX: translateX } }],
            { useNativeDriver: false }
        );

        const onHandlerStateChange = (event) => {
            if (event.nativeEvent.state === State.END) {
                const { translationX } = event.nativeEvent;

                if (translationX > 150) {
                    // Complete the swipe
                    Animated.timing(translateX, {
                        toValue: 200,
                        duration: 200,
                        useNativeDriver: false,
                    }).start(() => {
                        setIsCompleted(true);
                        setCurrentStatus('requested');
                    });
                } else {
                    // Reset the swipe
                    Animated.spring(translateX, {
                        toValue: 0,
                        useNativeDriver: false,
                    }).start();
                }
            }
        };

        return (
            <View style={styles.swipeContainer}>
                <View style={styles.swipeTrack}>
                    <Text style={styles.swipeText}>
                        {isCompleted ? 'Request Sent!' : 'Swipe to Join →'}
                    </Text>
                    <PanGestureHandler
                        onGestureEvent={onGestureEvent}
                        onHandlerStateChange={onHandlerStateChange}
                    >
                        <Animated.View
                            style={[
                                styles.swipeButton,
                                {
                                    transform: [{ translateX }],
                                    backgroundColor: isCompleted ? '#4CAF50' : '#FFD700',
                                },
                            ]}
                        >
                            <Ionicons
                                name={isCompleted ? 'checkmark' : 'arrow-forward'}
                                size={24}
                                color="#000"
                            />
                        </Animated.View>
                    </PanGestureHandler>
                </View>
            </View>
        );
    };

    const statusConfig = getStatusConfig(currentStatus);

    return (
        <SafeAreaView style={styles.container}>
            <StatusBar barStyle="light-content" backgroundColor="#1a1a1a" />
            <ScrollView contentContainerStyle={styles.scrollContainer}>
                {/* Group Image */}
                <Image source={{ uri: group.image }} style={styles.heroImage} />

                {/* Group Info */}
                <View style={styles.groupInfo}>
                    <Text style={styles.groupName}>{group.name}</Text>

                    {/* Mutual Friends */}
                    <View style={styles.mutualsRow}>
                        <View style={styles.mutualAvatars}>
                            {group.mutualFriends.slice(0, 4).map((avatar, index) => (
                                <Image
                                    key={index}
                                    source={{ uri: avatar }}
                                    style={[styles.mutualAvatar, { marginLeft: index > 0 ? -8 : 0 }]}
                                />
                            ))}
                        </View>
                    </View>

                    {/* Status Button */}
                    <TouchableOpacity
                        style={[styles.statusButton, { backgroundColor: statusConfig.bgColor }]}
                        onPress={handleStatusPress}
                    >
                        <Ionicons
                            name={statusConfig.icon}
                            size={20}
                            color="#fff"
                            style={styles.statusIcon}
                        />
                        <Text style={styles.statusText}>{statusConfig.text}</Text>
                    </TouchableOpacity>

                    {/* Distance and Tags Row */}
                    <View style={styles.metaRow}>
                        <View style={styles.distanceContainer}>
                            <Ionicons name="location-outline" size={16} color="#888" />
                            <Text style={styles.distance}>{group.distance}</Text>
                        </View>

                        <View style={styles.tagsContainer}>
                            {group.tags.map((tag, index) => (
                                <View key={index} style={styles.tag}>
                                    <Text style={styles.tagText}>{tag}</Text>
                                </View>
                            ))}
                        </View>
                    </View>

                    {/* Description */}
                    <View style={styles.descriptionSection}>
                        <Text style={styles.descriptionTitle}>Description</Text>
                        <Text style={styles.descriptionText}>{group.description}</Text>
                    </View>
                </View>
            </ScrollView>

            {/* Swipe to Join Button */}
            {currentStatus === 'join' && <SwipeToJoinButton />}
        </SafeAreaView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#1a1a1a',
    },
    scrollContainer: {
        paddingBottom: 120,
    },
    heroImage: {
        width: '100%',
        height: 250,
        resizeMode: 'cover',
    },
    groupInfo: {
        padding: 20,
    },
    groupName: {
        fontSize: 24,
        fontWeight: 'bold',
        color: '#fff',
        marginBottom: 16,
    },
    mutualsRow: {
        flexDirection: 'row',
        alignItems: 'center',
        marginBottom: 16,
    },
    mutualAvatars: {
        flexDirection: 'row',
        alignItems: 'center',
    },
    mutualAvatar: {
        width: 40,
        height: 40,
        borderRadius: 20,
        borderWidth: 2,
        borderColor: '#1a1a1a',
    },
    statusButton: {
        flexDirection: 'row',
        alignItems: 'center',
        justifyContent: 'center',
        paddingVertical: 12,
        paddingHorizontal: 24,
        borderRadius: 25,
        marginBottom: 20,
        alignSelf: 'flex-start',
    },
    statusIcon: {
        marginRight: 8,
    },
    statusText: {
        fontSize: 16,
        fontWeight: '600',
        color: '#fff',
    },
    metaRow: {
        marginBottom: 24,
    },
    distanceContainer: {
        flexDirection: 'row',
        alignItems: 'center',
        marginBottom: 12,
    },
    distance: {
        fontSize: 16,
        color: '#888',
        marginLeft: 6,
    },
    tagsContainer: {
        flexDirection: 'row',
        flexWrap: 'wrap',
    },
    tag: {
        backgroundColor: '#333',
        borderRadius: 20,
        paddingHorizontal: 16,
        paddingVertical: 8,
        marginRight: 10,
        marginBottom: 8,
    },
    tagText: {
        fontSize: 14,
        color: '#fff',
        fontWeight: '500',
    },
    descriptionSection: {
        marginTop: 8,
    },
    descriptionTitle: {
        fontSize: 18,
        fontWeight: '600',
        color: '#fff',
        marginBottom: 12,
    },
    descriptionText: {
        fontSize: 16,
        color: '#ccc',
        lineHeight: 24,
    },
    swipeContainer: {
        position: 'absolute',
        bottom: 30,
        left: 20,
        right: 20,
    },
    swipeTrack: {
        backgroundColor: '#333',
        height: 60,
        borderRadius: 30,
        justifyContent: 'center',
        alignItems: 'center',
        position: 'relative',
    },
    swipeText: {
        fontSize: 16,
        fontWeight: '600',
        color: '#fff',
    },
    swipeButton: {
        position: 'absolute',
        left: 5,
        width: 50,
        height: 50,
        borderRadius: 25,
        justifyContent: 'center',
        alignItems: 'center',
        elevation: 5,
        shadowColor: '#000',
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.3,
        shadowRadius: 4,
    },
});