import cv2
import numpy as np

Img = cv2.imread('./org.jpg')

# choose org image three point
pts1 = np.float32([[50,50],[200,50],[50,200]])
# choose new image three point
pts2 = np.float32([[10,100],[200,50],[100,250]])

M = cv2.getAffineTransform(pts1,pts2)
warpAffine1 = cv2.warpAffine(Img,M,dsize=(Img.shape[1]*2, Img.shape[0]*2))

status1 = cv2.imwrite('./warpAffine1.jpg',warpAffine1)


# choose new image three point
pts3 = np.float32([[10,300],[200,250],[200,350]])

M3 = cv2.getAffineTransform(pts1,pts3)
warpAffine3 = cv2.warpAffine(Img,M3,dsize=(Img.shape[1]*2, Img.shape[0]*2))

status3 = cv2.imwrite('./warpAffine3.jpg',warpAffine3)

# choose new image three point
pts4 = np.float32([[50,170],[200,150],[50,230]])

M4 = cv2.getAffineTransform(pts1,pts4)
warpAffine4 = cv2.warpAffine(Img,M4,dsize=(Img.shape[1], Img.shape[0]))

status4 = cv2.imwrite('./warpAffine4.jpg',warpAffine4)