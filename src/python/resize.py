import cv2
 
# read image as grey scale
Img = cv2.imread('./org.jpg')

# Scaling Bigger
ResizeImgB = cv2.resize(src=Img, dsize=(Img.shape[1]*2, Img.shape[0]*2))		  
	  
# save image
statusB = cv2.imwrite('./Scaling_big.jpg',ResizeImgB)

# Scaling Small

ImgC = cv2.resize(Img,(64,64), interpolation=cv2.INTER_CUBIC)

statusC = cv2.imwrite('./Scaling_small.jpg',ImgC)