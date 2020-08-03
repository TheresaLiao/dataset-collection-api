import cv2
import numpy as np

Img = cv2.imread('./org.jpg')

# x move 100px
# y move 50px
H = np.float32([[1,0,100],[0,1,50]])
Translation = cv2.warpAffine(Img,H,dsize=(Img.shape[1], Img.shape[0])) 

# save image
status1 = cv2.imwrite('./Translation.jpg',Translation)

# Rotation
RotateMatrix = cv2.getRotationMatrix2D(center=(Img.shape[1]/2, Img.shape[0]/2), angle=90, scale=1)
Rotation = cv2.warpAffine(Img, RotateMatrix,dsize=(Img.shape[1], Img.shape[0]))

# save image
status2 = cv2.imwrite('./Rotation.jpg',Rotation)