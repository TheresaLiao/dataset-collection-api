import cv2
import numpy as np

Img = cv2.imread('./org.jpg')

rows,cols,ch = Img.shape
pts1 = np.float32([[0,0],[0,cols],[rows,0],[rows,cols]])
pts2 = np.float32([[0,50],[0,cols-50],[rows,-50],[rows,cols+50]])

M = cv2.getPerspectiveTransform(pts1,pts2)
warpPerspective = cv2.warpPerspective(Img,M,dsize=(Img.shape[1]*2, Img.shape[0]*2))

status = cv2.imwrite('./warpPerspective.jpg',warpPerspective)

