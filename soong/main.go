package sm6150

import (
    "android/soong/android"
)

func init() {
    android.RegisterModuleType("meizu_sm6150_light_hal_binary", lightHalBinaryFactory)
}
